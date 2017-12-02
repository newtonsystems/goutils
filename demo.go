package main

func main() {
	logger := NewServiceLogger("new-service")

	logger.Debug("msg", "This is a debug message")
	logger.Info("msg", "This is a info message")
	logger.Stage("msg", "This is a stage message")
	logger.Warn("msg", "This is a warn message")
	logger.Err("msg", "This is a error message")
	logger.Crit("msg", "This is a critical message (we add an extra empty line)")

	println("")
	println("We also provide convenient functions for just simple string log messages: \t (adds message to 'msg' key)")
	logger.DebugMsg("This is a demonstration of the DebugMsg().")
	logger.InfoMsg("This is a demonstration of the InfoMsg().")
	logger.StageMsg("This is a demonstration of the StageMsg().")
	logger.WarnMsg("This is a demonstration of the WarnMsg().")
	logger.ErrMsg("This is a demonstration of the ErrMsg().")
	logger.CritMsg("This is a demonstration of the CritMsg() (we add an extra empty line).")

	println("")
	println("We also provide convenient functions for just simple string log messages: \t (adds message to 'msg' key)")
	logWith := logger.With("component", "with")
	logWith.Debug("msg", "This is a debug message with a component (With())")
	logWith.DebugMsg("This is a demonstration of the DebugMsg() with a component (With())")

	println("")
	println("You can also have prefixes:")
	logPrefix := logger.WithPrefix("request", "[f637463:demo]")
	logPrefix.Debug("msg", "This is a debug message - WithPrefix()")
	println("")

}
