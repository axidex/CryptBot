package telegram

type Command string

const (
	CommandStart   Command = "/start"
	CommandHelp    Command = "/help"
	CommandElGamal Command = "/elgamal"
	CommandDiffie  Command = "/diffie"
	CommandRSA     Command = "/rsa"
	CommandEnigma  Command = "/enigma"
	CommandDes     Command = "/des"
)

func (c Command) GetHint() string {
	switch c {
	case CommandStart:
		return StartHint
	case CommandHelp:
		return StartHint
	case CommandElGamal:
		return ElGamalHint
	case CommandDiffie:
		return DiffieHellmanHint
	case CommandRSA:
		return RSAHint
	case CommandEnigma:
		return EnigmaHint
	case CommandDes:
		return DesHint
	}

	return ""
}

func GetCommand(command string) Command {
	switch command {
	case "/start":
		return CommandStart
	case "/help":
		return CommandHelp
	case "/elgamal":
		return CommandElGamal
	case "/diffie":
		return CommandDiffie
	case "/rsa":
		return CommandRSA
	case "/enigma":
		return CommandEnigma
	case "/des":
		return CommandDes
	}
	return CommandHelp
}
