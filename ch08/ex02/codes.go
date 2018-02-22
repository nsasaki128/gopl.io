package main

//see https://tools.ietf.org/rfc/rfc640.txt
const (
	CommandOk = 200
	//CommandSyntaxError                = 500
	//ParamsSyntaxError                 = 501
	//CommandNotImplementedSuperfluous  = 202
	CommandNotImplemented = 502
	//BadSequenceCommand                = 503
	CommandNotImplementedForParameter = 504
	//RestartMarkerReply                = 110
	//SystemStatus                      = 211
	//DirectoryStatus                   = 212
	//FileStatus                        = 213
	//HelpMessage                       = 214
	//ServiceReadyInNMinutes            = 120
	ServiceReadyForNewUser = 220
	//ServiceClosingTELNETConnection    = 221
	//ServiceNotAvailable               = 421
	//DataConnectionAlreadyOpen         = 125
	//DataConnectionOpen                = 225
	CantOpenDataConnection = 425
	//ClosingDataConnection             = 226
	ConnectionTrouble = 426
	Entering          = 227
	UserLoggedOn      = 230
	//NotLoggedIn                       = 530
	//UserNameOkay                      = 331
	//NeedAccountForLogin               = 332
	//NeedAccountForStoringFiles        = 532
	FileStatusOkay          = 150
	RequestedFileActionOkey = 250
	//RequestedFileActionPending        = 350
	FileUnavailable     = 450
	FileUnavailableBusy = 550
	//LocalErrorInProcessing            = 451
	//InsufficientStorageSpace          = 452
	//ExceededStorageAllocation         = 552
	//FileNameNotAllowed                = 553
	//StartMailInput                    = 354
	PathNameCreated = 257
)
