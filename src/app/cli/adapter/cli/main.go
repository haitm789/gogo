package cli

type (
	Command struct {
		Command string
		Caption string
		Desc    string
		Func    func(args []string)
	}
)
