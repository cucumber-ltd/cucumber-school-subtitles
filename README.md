# cucumber-school-captions

Generating captions for Cucumber School videos!

## Steps

* Use [autosub](https://github.com/agermanidis/autosub) for the initial generation of VTT files.
  * Example: `autosub -F vtt 02_your_first_scenario/java/Cucumber\ School\ -\ 01x02\ -\ Your\ First\ Scenario\ \(Java\).mp4`
* Autosub uses the Google Web Speech API to generate the test in the files and it's always going to need cleaning up. Expect misspellings, incorrect words, missing lines, and timing issues to all be present. There are a couple different ways clean up can be done, but what I've started doing is:
  * Run through the video with the caption file open to fix things like misspelled/incorrect words and missing dialogue without worrying about timing
  * Run through again to fix timing, formatting, etc. Keep in mind the style guidelines below when doing this.
* Use ffmpeg to combine the caption file with the video
  * Command tbd
* Upload
  * Location tbd
* Success!

## Style Guidelines

Some notes taken from a couple different [style guidelines](http://bbc.github.io/subtitle-guidelines/) and decided on by myself.

* No ending periods.
* The first letter of every frame must be capitalized, regardless of where in a sentence the frame begins.
* BDD, Cucumber, Gherkin, Three Amigos, and any other named items must be capitalized
* Maximum of 2 lines per caption frame
* Maximum of 40 characters per line
* Minimum of 1.5 seconds of on screen time for shorter frames
* Based on the recommended rate of 160-180 words per minute, you should aim to leave a subtitle on screen for a minimum period of around 0.3 seconds per word (e.g. 1.2 seconds for a 4-word subtitle). This isn't a hard and fast rule, but if the caption frame doesn't seem like it's on the screen long enough to be easily read, try adjusting the times.