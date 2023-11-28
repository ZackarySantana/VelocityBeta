package jobs

type Job interface {
	SetupCommand() []string
	GetImage() string
	GetCommand() string
	GetName() string
}

type CommandJob struct {
	Image string

	Command string
	Name    string
}

func (j *CommandJob) SetupCommand() []string {
	return []string{}
}

func (j *CommandJob) GetImage() string {
	return j.Image
}

func (j *CommandJob) GetCommand() string {
	return j.Command
}

func (j *CommandJob) GetName() string {
	return j.Name
}

type FrameworkJob struct {
	Image string

	Language  string
	Framework string
	Name      string
}

func (j *FrameworkJob) SetupCommand() []string {
	return []string{}
}

func (j *FrameworkJob) GetImage() string {
	return j.Image
}

func (j *FrameworkJob) GetCommand() string {
	return "echo hello world"
}

func (j *FrameworkJob) GetName() string {
	return j.Name
}
