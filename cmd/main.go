package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/saipulmuiz/go-project-starter/config"
)

var (
	signalShutdown = []os.Signal{ // https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
		os.Interrupt,
		syscall.SIGTERM, // The SIGTERM signal is a generic signal used to cause program termination.
		syscall.SIGINT,  // The SIGINT (“program interrupt”) signal is sent when the user types the INTR character (normally C-c)
		syscall.SIGKILL, // The SIGKILL signal is used to cause immediate program termination
		syscall.SIGHUP,  // The SIGHUP (“hang-up”) signal is used to report that the user’s terminal is disconnected
	}
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer os.Exit(0)

	var (
		err     error
		timeout = 30 * time.Second
	)

	exit := func(ctx context.Context, code int, msg ...string) {
		if len(msg) > 0 {
			log.Printf("INFO: %s", msg)
		}

		<-config.GlobalShutdown.GracefullyShutdown(ctx, timeout)
		log.Printf("INFO: the application successfully to shuting down, bye...")
		os.Exit(code)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		ec = make(chan error, 1)
		// done = make(chan struct{}, 1)
	)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, signalShutdown...)

	select {
	case err = <-ec:
		log.Printf("ERROR: got error during start the service %s", err.Error())
		break
	case sign := <-terminateSignals:
		cancel()
		log.Print("INFO: received termination signal %s from operation system\n", sign)
		break
	}

	exit(context.Background(), 0)
}
