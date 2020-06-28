package version

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	GITCOMMIT     string
	GITCOMMITTIME string
	GITTAG        string
	PRONAME       string
)

func MockProName() string {
	if PRONAME == "" {
		PRONAME = "foo-system-bar-login"
	}
	return PRONAME
}

func Version() string {
	var s = strings.Builder{}
	s.WriteString(fmt.Sprintf("%5v: %s\n", "GIT-COMMIT", GITCOMMIT))
	s.WriteString(fmt.Sprintf("%5v: %s\n", "GIT-COMMIT-TIME", GITCOMMITTIME))
	s.WriteString(fmt.Sprintf("%5v: %s\n", "GIT-TAG", GITTAG))
	s.WriteString(fmt.Sprintf("%5v: %s\n", "PRO-NAME", PRONAME))
	return s.String()
}

func ShortGitCommit() string {
	if len(GITCOMMIT) >= 8 {
		return GITCOMMIT[:8]
	}
	return GITCOMMIT
}

func LoadversionCmd(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version())
		}})
}
