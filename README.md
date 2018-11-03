# cucumber-school-captions

Generating captions for Cucumber School videos!

## Vocabulary

* **caption frame** - a set of text displayed between a start and end time in a video. Example:

```text
2
00:00:19.712 --> 00:00:23.900
Last time we learned how to use
conversations about concrete examples
```

## Steps

* Use [autosub](https://github.com/agermanidis/autosub) for the initial generation of VTT files.
  * Example: `autosub -F vtt 02_your_first_scenario/java/Cucumber\ School\ -\ 01x02\ -\ Your\ First\ Scenario\ \(Java\).mp4`
* Autosub uses the Google Web Speech API to generate the test in the files and it's always going to need cleaning up. Expect misspellings, incorrect words, missing lines, and timing issues to all be present. There are a couple different ways clean up can be done, but what I've started doing is:
  * Using VLC, run through the video with the caption file open to fix things like misspelled/incorrect words and missing dialogue without worrying about timing
  * Run through again to fix timing, formatting, etc. Keep in mind the style guidelines below when doing this.
  * To fix caption frame numbering for files run `make renumber`. To run for a specific file, prefix with a FILENAME value to the file you'd like to change. Ex: `FILENAME="./02_your_first_scenario/java/Cucumber School - 01x02 - Your First Scenario (Java).vtt" make renumber`
* Use ffmpeg to combine the caption file with the video
  * Example: `ffmpeg -i "Cucumber School - 01x01 - Give Me An Example.mp4" -i "Cucumber School - 01x01 - Give Me An Example.vtt" -metadata:s:s:0 language=eng -c copy -c:s mov_text "Cucumber School - 01x01 - Give Me An Example - Captions.mp4"`
* Upload
  * Location & process TBA
* Success!

## Style Guidelines

Some notes taken from a couple different [style guidelines](http://bbc.github.io/subtitle-guidelines/) and decided on by myself.

* No ending periods.
* Prefer to have one caption lead immediately into another where it makes sense. This is meant to make the caption reading experience less jarring and smoother. The start time of one frame should be the same as the frame that precedes it so that there is no 'flashing' as one caption disappears and the next appears. Cases when and when not to do this follow:
  * Sentences that span multiple caption frames must be connected
  * Sentences that are within approximately a half second of one another should also be connected.
* BDD, Cucumber, Gherkin, Three Amigos, and any other named items must be capitalized
* Maximum of 2 lines per caption frame
* Maximum of 50 characters per line, where possible, with preference toward shorter lines
* Minimum of 1 seconds of on screen time for shorter frames if possible
* For commands that are to be input, wrap them in 'single quotes'
* For file names or other things like variables, if they are capitalized in the video, make them capital in the caption.

## Status

Captions and videos go through a number of statuses in their life cycle and will be tracked below.

* Initial Run - Running through the base autosub output to fix words that are wrong, missing, or phrases that are missing. This is meant to get the captions to match the script's wording without the primary focus being timing or lines.
* Main Run - This step's goal is to get the caption frames timed correctly and broken out into ways that flow well. `make renumber` should be run at the end of this step.
* Polish Run - This is the final step for the caption files, meant to ensure that timings and flow are correct before being added to the video.
* Captions in Video - This step is when `make embed [path]` should be run with the path being the directory the video and the caption file are located in. **TODO: Write embed step**
* Complete - 🎉

| Video | Language | Initial Run | Main Run | Polish Run | Captions in Video | Complete |
|:-----:|:--------:|:-----------:|:--------:|:----------:|-------------------|----------|
| 01    |          | √           | √        |            | √                 |          |
| 02    |          |             |          |            |                   |          |
|       | Ruby     | √           | √        |            |                   |          |
|       | Java     | √           | √        |            |                   |          |
| 03    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 04    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 05    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 06    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 07    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 08    |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 09-10 |          |             |          |            |                   |          |
|       | Ruby     |             |          |            |                   |          |
|       | Java     |             |          |            |                   |          |
| 11-12 |          |             |          |            |                   |          |
|       | Ruby     | √           | √        |            |                   |          |
|       | Java     |             |          |            |                   |          |s
