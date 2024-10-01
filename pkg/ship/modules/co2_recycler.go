package modules

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/bad-noodles/ss-lovelace/pkg/message"
	"github.com/bad-noodles/ss-lovelace/pkg/theme"
)

var (
	carbon = 12.011
	oxygen = 15.999
)

var undesireds = []string{
	fmt.Sprintf("%f|%f|%f", carbon, oxygen, oxygen),
	fmt.Sprintf("%f|%f", carbon+oxygen, oxygen),
}
var desired = fmt.Sprintf("%f|%f", carbon, oxygen*2)

func undesiredChain() (int, string) {
	length := rand.IntN(100)
	var output strings.Builder
	for i := 0; i < length; i++ {
		u := rand.IntN(len(undesireds))
		output.WriteString(undesireds[u])
		output.WriteString("-")
	}

	return length, output.String()
}

func desiredChain() (int, string) {
	length := rand.IntN(50)
	if length == 0 {
		length = 1
	}
	var output strings.Builder
	for i := 0; i < length; i++ {
		output.WriteString(desired)
		output.WriteString("-")
	}

	return length, output.String()
}

type Co2Recycler struct {
	expectedAnswer string
}

func (r *Co2Recycler) Name() string {
	return "CO₂ Recycler"
}

func (r *Co2Recycler) SendChallenge() string {
	lc1, c1 := undesiredChain()
	lc2, c2 := desiredChain()
	_, c3 := undesiredChain()
	challenge := fmt.Sprintf("%s%s%s", c1, c2, c3)
	challenge = challenge[:len(challenge)-1]

	r.expectedAnswer = fmt.Sprintf("%d|%d", lc1, lc2)

	return challenge
}

func (r *Co2Recycler) ValidateChallenge(answer string) bool {
	return answer == r.expectedAnswer
}

func (r *Co2Recycler) Message() message.Message {
	return message.Message{
		Subject: "CO₂ Recycler ♻️",
		Body: `Recycling the CO₂ allows us to recapture part of the oxygen we would usually discard.

It is a very important module to be operational, as otherwise we will go through our oxygen reserves too fast!

The problem is that not everything that comes out of the recycler is good for us to use, so we need to filter it out and keep only the good stuff.

The recycler uses high-speed collision to break apart the CO₂, but that can have three different outcomes:
  - C + O + O
  - CO + O
  - C + O₂

The first two outcomes are not interesting for us, we only care about the outcome that carries an O₂, as that is what we need for breathing.

We have a ` + theme.Bold.Render("sliding window") + ` filter that should be able to detect the correct result based on their weight, but that is not working right now!

Taking into consideration the information below, please fix the filter:

The weight of a carbon atom is 12.011 and of an oxygen atom is 15.999.

The weight of each atom or molecule will be separated by a '|' character, so 'C + O + O' looks like '12.011|15.999|15.999'

The recycler will send multiple outcomes in one go, separated by a '-' character.

The correct outcome comes out clumped together, so we only need the index of the first occurence and how many are there.

That will cause the filter to slide to the correct position and filter out the bad stuff.

The response should follow this pattern: '{index}|{length}\n'


Be careful! Bugs in this logic might cause unbreathable gas to get into the system, causing people to suffocate and die!







    


`,
	}
}
