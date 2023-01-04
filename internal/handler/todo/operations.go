package todo

type Operation uint

const (
	List Operation = iota
	Add
	Del
	Show
	Edit
)

func (o Operation) String() string {
	switch o {
	case List:
		return "/todos"
	case Add:
		return "/todos/add"
	case Del:
		return "/todos/del"
	case Show:
		return "/todos/show"
	case Edit:
		return "/todos/edit"
	default:
		return "/"
	}
}
