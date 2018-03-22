package main

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	//元のnotifyUserを保存しておいて回復する
	saved := notifyUser
	defer func() { notifyUser = saved }()
	//テストのため偽のnotifyUserを設定する
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	//...980MBが使われた状態を装う...

	const user = "joe@example.org"
	usage[user] = 980000000 // simulate a 980MB-used condition

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifiedUser not called")
	}
	if notifiedUser != user {
		t.Errorf("Wrong user (%s) notified, want %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
