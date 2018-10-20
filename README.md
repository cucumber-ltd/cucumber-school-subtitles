# cucumber-school-captions

Generating captions for Cucumber School videos!

## Steps

* Use [autosub](https://github.com/agermanidis/autosub) for the initial generation of VTT files. Clean up will need to be done by hand.
  * Example: `autosub -F vtt 02_your_first_scenario/java/Cucumber\ School\ -\ 01x02\ -\ Your\ First\ Scenario\ \(Java\).mp4`
* Run through video with caption file to clean up the inevitable errors. Misspellings, incorrect words, missing lines, and timing issues may all pop up.
* Use ffmpeg to combine the caption file with the video
* Upload
* Success!

## Style Guideline

Some notes taken from a couple different [style guidelines](http://bbc.github.io/subtitle-guidelines/) and decided on by myself.

* No punctuation except for question marks or exclamation points.
* No more than 2 lines per caption frame
* Maximum of 37 characters per line
* Minimum of 1.5 seconds of on screen time for shorter frames
* BDD, Cucumber, and Gherkin must be capitalized
* Based on the recommended rate of 160-180 words per minute, you should aim to leave a subtitle on screen for a minimum period of around 0.3 seconds per word (e.g. 1.2 seconds for a 4-word subtitle).