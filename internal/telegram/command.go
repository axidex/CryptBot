package telegram

type Command string

const (
	CommandStart    Command = "/start"
	CommandHelp     Command = "/help"
	CommandElGamal  Command = "/elgamal"
	CommandDiffie   Command = "/diffie"
	CommandRSA      Command = "/rsa"
	CommandEnigma   Command = "/enigma"
	CommandDes      Command = "/des"
	CommandAes      Command = "/aes"
	CommandInvMix   Command = "/invMix"
	CommandFeistel  Command = "/feistel"
	CommandHash     Command = "/hash"
	CommandA5       Command = "/A5"
	Command3Des     Command = "/3des"
	CommandBlowFish Command = "/blowfish"
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
	case CommandAes:
		return AesHint
	case CommandInvMix:
		return InvMixHint
	case CommandFeistel:
		return FeistelHint
	case CommandHash:
		return HashHint
	case CommandA5:
		return A5Hint
	case Command3Des:
		return Des3Hint
	case CommandBlowFish:
		return BlowFishHint
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
	case "/aes":
		return CommandAes
	case "/invmix":
		return CommandInvMix
	case "/feistel":
		return CommandFeistel
	case "/hash":
		return CommandHash
	case "/a5":
		return CommandA5
	case "/3des":
		return Command3Des
	case "/blowfish":
		return CommandBlowFish
	}
	return CommandHelp
}
