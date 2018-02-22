package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

type ftp struct {
	conn net.Conn
	cwd  string
	addr string
}

func newFtp(conn net.Conn) *ftp {
	return &ftp{
		conn: conn,
		cwd:  "/",
		addr: conn.RemoteAddr().String(),
	}
}

func (ftp *ftp) reply(code int, text string) {
	io.WriteString(ftp.conn, fmt.Sprintf("%d %s\r\n", code, text))
}

func (ftp *ftp) run() {
	defer ftp.conn.Close()
	ftp.reply(ServiceReadyForNewUser, "Service ready for new user.")

	scanner := bufio.NewScanner(ftp.conn)
	for scanner.Scan() {
		command := scanner.Text()
		log.Println(command)
		if !ftp.handleConn(command) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error:", err)
	}
}

func (ftp *ftp) handleConn(c string) bool {
	command := strings.Split(c, " ")
	switch command[0] {
	case "USER":
		ftp.handleUserCommand(c)
	case "TYPE":
		ftp.handleTypeCommand(c)
	case "PASV":
		ftp.handlePasvCommand(c)
	case "CWD":
		ftp.handleCwdCommand(c)
	case "LIST":
		ftp.handleListCommand(c)
	case "STOR":
		ftp.handleStorCommand(c)
	case "SYST":
		ftp.handleSystCommand(c)
	case "PWD":
		ftp.handlePwdCommand(c)
	case "QUIT":
		ftp.handleQuitCommand(c)
		return false
	default:
		ftp.reply(CommandNotImplemented, "Command not implemented.")
	}
	return true
}

func (ftp *ftp) handleUserCommand(c string) {
	ftp.reply(UserLoggedOn, "User Logged on.")
}

func (ftp *ftp) handleTypeCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}

	typeCode := command[1]
	if typeCode != "A" && typeCode != "I" {
		ftp.reply(CommandNotImplementedForParameter, "Command not implemented for that parameter.")
		return
	}
	ftp.reply(CommandOk, "Command okay.")
}
func (ftp *ftp) handlePasvCommand(c string) {
	listenIP := ftp.addr
	lastIdx := strings.LastIndex(listenIP, ":")
	socket, _ := strconv.Atoi(listenIP[lastIdx+1:])
	quads := strings.Split(listenIP[:lastIdx], ".")

	p1 := socket / 256
	p2 := socket - (p1 * 256)
	target := fmt.Sprintf("(%s,%s,%s,%s,%d,%d)", quads[0], quads[1], quads[2], quads[3], p1, p2)
	msg := "Entering Passive Mode " + target
	fmt.Println(msg)
	ftp.reply(Entering, msg)
}

func (ftp *ftp) handleCwdCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}
	cm := command[1]

	if !strings.HasPrefix(cm, "/") {
		cm = path.Join(ftp.cwd, cm)
	}

	info, err := os.Stat(cm)
	if err != nil || !info.IsDir() {
		ftp.reply(FileUnavailableBusy, "Directory not found.")
		return
	}
	fmt.Println(info.Name())
	ftp.cwd = cm

	ftp.reply(RequestedFileActionOkey, "Requested file action okay, completed.")
}
func (ftp *ftp) handleListCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}
	p := path.Join(ftp.cwd, command[1])

	ftp.reply(FileStatusOkay, "Open data connection.")
	conn, err := net.Dial("tcp", ftp.addr)
	if err != nil {
		log.Println("error: ", err)
		ftp.reply(CantOpenDataConnection, "Can't open data connection.")
		return
	}
	defer conn.Close()

	info, err := os.Stat(p)
	if err != nil {
		ftp.reply(FileUnavailableBusy, "File or Directory not found.")
		return
	}

	var result []byte
	if info.IsDir() {
		is, err := ioutil.ReadDir(p)
		if err != nil {
			ftp.reply(FileUnavailableBusy, "File or Directory not found.")
			return
		}

		for _, i := range is {
			result = append(result, []byte(i.Name()+"\r\n")...)
		}
	} else {
		result = []byte(fmt.Sprintf("%s\r\n", p))
	}

	if _, err = conn.Write(result); err != nil {
		log.Println("error: ", err)
		ftp.reply(ConnectionTrouble, "Connection closed; transfer aborted.")
		return
	}

	ftp.reply(RequestedFileActionOkey, "File transfer completed.")
}

func (ftp *ftp) handleStorCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}
	p := path.Join(ftp.cwd, command[1])

	ftp.reply(FileStatusOkay, "File status okay; about to open data connection.")

	conn, err := net.Dial("tcp", ftp.addr)
	if err != nil {
		log.Println("error: ", err)
		ftp.reply(CantOpenDataConnection, "Can't open data connection.")
		return
	}
	defer conn.Close()

	f, err := os.Create(p)
	if err != nil {
		log.Println("error: ", err)
		ftp.reply(FileUnavailable, "Requested file action not taken.")
		return
	}
	defer f.Close()

	if _, err = io.Copy(f, conn); err != nil {
		log.Println("error: ", err)
		ftp.reply(FileUnavailable, "Requested file action not taken.")
		return
	}

	ftp.reply(RequestedFileActionOkey, "File transfer completed.")
}
func (ftp *ftp) handleSystCommand(c string) {
	ftp.reply(CommandOk, "Simple FTP")
}

func (ftp *ftp) handlePwdCommand(c string) {
	fmt.Println(ftp.cwd)
	ftp.reply(PathNameCreated, ftp.cwd)
}
func (ftp *ftp) handleQuitCommand(c string) {
	ftp.reply(ServiceClosingTELNETConnection, "Close connection.")
}
