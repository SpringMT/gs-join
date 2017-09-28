package main

import(
	"fmt"
	"runtime"
	"os"

	"github.com/urfave/cli"
	"bufio"
	"strings"
)

var exitCode = 0

func main()  {
	app := cli.NewApp()
	app.Name = "gs-join"
	app.Version = "1"

	app.Usage = `test
	test
	`
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "delimiter, d",
			Usage: "Use character char as a field delimiter",
		},
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(app.Writer, `%v version %v
Compiler: %s %s
`,
			app.Name,
			app.Version,
			runtime.Compiler,
			runtime.Version())
	}
	app.Action = doMain
	app.Run(os.Args)
	os.Exit(exitCode)
}

func doMain(c *cli.Context) error {
	words := make([]string, 0, 0)
	delimiter := c.String("delimiter")

	if len(c.Args()) > 0 {
		for _, filename := range c.Args() {
			file, err := os.Open(filename)
			if err != nil {
				os.Stderr.WriteString("failed to open and read `" + filename + "`.\n")
				exitCode = 1
				return err
			}
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()
				words = append(words, line)
			}
			file.Close()
			if err != nil {
				os.Stderr.WriteString("reading input error\n")
				exitCode = 1
				return err
			}
		}
	} else {
		file := os.Stdin
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			words = append(words, line)
		}
		file.Close()
		if scanner.Err() != nil {
			os.Stderr.WriteString("reading input error\n")
			exitCode = 1
			return scanner.Err()
		}
	}

	println(strings.Join(words, delimiter))
	return nil
}
