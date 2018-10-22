# cucumber-school-captions

Generating captions for Cucumber School videos!

## Steps

* Use [autosub](https://github.com/agermanidis/autosub) for the initial generation of VTT files.
  * Example: `autosub -F vtt 02_your_first_scenario/java/Cucumber\ School\ -\ 01x02\ -\ Your\ First\ Scenario\ \(Java\).mp4`
* Autosub uses the Google Web Speech API to generate the test in the files and it's always going to need cleaning up. Expect misspellings, incorrect words, missing lines, and timing issues to all be present. There are a couple different ways clean up can be done, but what I've started doing is:
  * Using VLC, run through the video with the caption file open to fix things like misspelled/incorrect words and missing dialogue without worrying about timing
  * Run through again to fix timing, formatting, etc. Keep in mind the style guidelines below when doing this.
* Use ffmpeg to combine the caption file with the video
  * Example: `ffmpeg -i "Cucumber School - 01x01 - Give Me An Example.mp4" -i "Cucumber School - 01x01 - Give Me An Example.vtt" -metadata:s:s:0 language=eng -c copy -c:s mov_text "Cucumber School - 01x01 - Give Me An Example - Captions.mp4"`
* Upload
  * Location & process TBA
* Success!

## Style Guidelines

Some notes taken from a couple different [style guidelines](http://bbc.github.io/subtitle-guidelines/) and decided on by myself.

* No ending periods.
* For sentences that span multiple caption frames, butt the timing of the ending and beginning of each frame together for a smoother experience. For example: Frame 1 ends around 00:00:10.575 and extends across frames, so frame 2 should start at the same time.
* BDD, Cucumber, Gherkin, Three Amigos, and any other named items must be capitalized
* Maximum of 2 lines per caption frame
* Maximum of 40 characters per line (where possible)
* Minimum of 1.5 seconds of on screen time for shorter frames
* Based on the recommended rate of 160-180 words per minute, you should aim to leave a subtitle on screen for a minimum period of around 0.3 seconds per word (e.g. 1.2 seconds for a 4-word subtitle). This isn't a hard and fast rule, but if the caption frame doesn't seem like it's on the screen long enough to be easily read, try adjusting the times.