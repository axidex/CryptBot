package problems

import (
	"errors"
	"fmt"
	"strings"
)

type Rotor struct {
	Encoder map[rune]rune
	Decoder map[rune]rune
	Name    string
}

func (c *Rotor) Encode(char rune) rune {
	return c.Encoder[char]
}

func (c *Rotor) Decode(char rune) rune {
	return c.Decoder[char]
}

type Reflector struct {
	Coder map[rune]rune
	Name  string
}

func (c *Reflector) Code(char rune) rune {
	return c.Coder[char]
}

type PatchPanel struct {
	Coder map[rune]rune
}

func (c *PatchPanel) Code(char rune) rune {
	newChar, ok := c.Coder[char]
	if !ok {
		return char
	}
	return newChar
}

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r1       = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	r2       = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	r3       = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	r4       = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
	r5       = "VZBRGITYUPSDNHLXAWMJQOFECK"
	r6       = "JPGVOUMFYQBENHZRDKASXLICTW"
	r7       = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
	r8       = "FKQHTLXOCBJSPDZRAMEWNIUYGV"
	refB     = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	refC     = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
)

var rotors = []string{r1, r2, r3, r4, r5, r6, r7, r8}

func GetRotor(rotorNumber int) (*Rotor, error) {
	name := fmt.Sprintf("r%d", rotorNumber)
	rotorNumber--
	if rotorNumber > len(rotors) || rotorNumber < 0 {
		return nil, errors.New("rotor number out of range")
	}
	rotor := rotors[rotorNumber]
	encoder := createAlphabetReference(alphabet, rotor)
	decoder := createAlphabetReference(rotor, alphabet)

	return &Rotor{Encoder: encoder, Decoder: decoder, Name: name}, nil
}

func GetReflector(reflector string) (*Reflector, error) {
	switch strings.ToLower(reflector) {
	case "b":
		return &Reflector{Coder: createAlphabetReference(alphabet, refB), Name: "refB"}, nil
	case "c":
		return &Reflector{Coder: createAlphabetReference(alphabet, refC), Name: "refC"}, nil
	default:
		return nil, errors.New("unknown reflector " + reflector)
	}
}

func GetPatchPanel(patchPanel string) (*PatchPanel, error) {
	coder := make(map[rune]rune)
	pairs := strings.Split(patchPanel, ",")
	for _, pair := range pairs {
		if len(pair) != 3 || pair[1] != '-' {
			return nil, fmt.Errorf("invalid pair format: %s", pair)
		}

		a := rune(pair[0])
		b := rune(pair[2])

		coder[a] = b
		coder[b] = a
	}

	return &PatchPanel{Coder: coder}, nil
}

func createAlphabetReference(alphabet1, alphabet2 string) map[rune]rune {
	mapping := make(map[rune]rune)
	for i, char := range alphabet1 {
		mapping[char] = rune(alphabet2[i])
	}
	return mapping
}

func Enigma(rotorNumbers int, reflectorName, phrase, patchPanel string) (string, error) {
	var currentRotors []Rotor
	rotorDigits := ConvertToDigits(rotorNumbers)

	for _, rotorDigit := range rotorDigits {
		rotor, err := GetRotor(rotorDigit)
		if err != nil {
			return "", err
		}
		currentRotors = append(currentRotors, *rotor)
	}

	reflector, err := GetReflector(reflectorName)
	if err != nil {
		return "", err
	}

	pp, err := GetPatchPanel(patchPanel)
	if err != nil {
		return "", err
	}

	pipeline := CreatePipeline(currentRotors, *pp, *reflector)

	var newRunes []rune
	var solutions []string
	for idx, char := range phrase {
		newRune, solution := pipeline.Run(idx, char)
		newRunes = append(newRunes, newRune)
		solutions = append(solutions, solution)
	}

	taskSolution := strings.Join(solutions, "\n\n")

	return strings.Join([]string{taskSolution, string(newRunes)}, "\n\n"), nil
}

type Pipeline struct {
	pp     PatchPanel
	rotors []Rotor
	ref    Reflector
}

func CreatePipeline(rotors []Rotor, pp PatchPanel, ref Reflector) *Pipeline {
	return &Pipeline{
		pp:     pp,
		rotors: rotors,
		ref:    ref,
	}
}

func (p *Pipeline) Title(idx int) string {
	var text []string
	text = append(text, "I")
	text = append(text, "PP")
	text = append(text, fmt.Sprintf("+%d", idx))
	text = append(text, p.rotors[0].Name)
	text = append(text, fmt.Sprintf("-%d", idx))
	text = append(text, p.rotors[1].Name)
	text = append(text, p.rotors[2].Name)
	text = append(text, p.ref.Name)
	text = append(text, p.rotors[2].Name)
	text = append(text, p.rotors[1].Name)
	text = append(text, fmt.Sprintf("-%d", idx))
	text = append(text, p.rotors[0].Name)
	text = append(text, fmt.Sprintf("+%d", idx))
	text = append(text, "PP")
	text = append(text, "E")

	return strings.Join(text, "->")

}

func (p *Pipeline) Run(idx int, rune rune) (rune, string) {
	var text []string
	toAdd := idx + 1

	text = append(text, string(rune))

	//text = append(text, "PP")
	rune = p.pp.Code(rune)
	text = append(text, string(rune))

	//text = append(text, "ADD")
	rune = AddRuneToLetter(rune, toAdd)
	text = append(text, string(rune))

	//text = append(text, "First rotor encode")
	rune = p.rotors[0].Encode(rune)
	text = append(text, string(rune))

	//text = append(text, "SUB")
	rune = SubtractRuneFromLetter(rune, toAdd)
	text = append(text, string(rune))

	//text = append(text, "Second rotor encode")
	rune = p.rotors[1].Encode(rune)
	text = append(text, string(rune))

	//text = append(text, "Third rotor encode")
	rune = p.rotors[2].Encode(rune)
	text = append(text, string(rune))

	//text = append(text, "Reflector")
	rune = p.ref.Code(rune)
	text = append(text, string(rune))

	//text = append(text, "Third rotor decode")
	rune = p.rotors[2].Decode(rune)
	text = append(text, string(rune))

	//text = append(text, "Second rotor decode")
	rune = p.rotors[1].Decode(rune)
	text = append(text, string(rune))

	//text = append(text, "SUB")
	rune = AddRuneToLetter(rune, toAdd)
	text = append(text, string(rune))

	//text = append(text, "First rotor decode")
	rune = p.rotors[0].Decode(rune)
	text = append(text, string(rune))

	//text = append(text, "ADD")
	rune = SubtractRuneFromLetter(rune, toAdd)
	text = append(text, string(rune))

	//text = append(text, "PP")
	rune = p.pp.Code(rune)
	text = append(text, string(rune))

	return rune, strings.Join([]string{p.Title(toAdd), strings.Join(text, " -> ")}, "\n")
}
