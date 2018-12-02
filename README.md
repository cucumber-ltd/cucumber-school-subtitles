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

* Upload video as unlisted to Youtube. (YT) Their service should automatically generate a subtitles file.
* Edit YT's file in their interface to correct anything that might be wrong with the auto generated text.
* Update the caption file and then download it in its `.srt` format.
* Use ffmpeg to combine the caption file with the video
  * Example: `ffmpeg -i "Cucumber School - 01x01 - Give Me An Example.mp4" -i "Cucumber School - 01x01 - Give Me An Example.vtt" -metadata:s:s:0 language=eng -c copy -c:s mov_text "Cucumber School - 01x01 - Give Me An Example - Captions.mp4"`
* Upload
  * Location & process TBA
* Success!

## Style Guidelines

Some notes taken from a couple different [style guidelines](http://bbc.github.io/subtitle-guidelines/) and decided on by myself.

* No ending commas or periods
* Prefer to have one caption lead immediately into another where it makes sense. This is meant to make the caption reading experience less jarring and smoother. The start time of one frame should be the same as the frame that precedes it so that there is no 'flashing' as one caption disappears and the next appears. Cases when and when not to do this follow:
  * Sentences that span multiple caption frames must be connected
  * Sentences that are within approximately a half second of one another should also be connected.
* Maximum of 2 lines per caption frame
* Maximum of 45 characters per line, where possible, with preference toward shorter lines
* Minimum of 1 seconds of on screen time for shorter frames if possible
* For commands that are to be input, wrap them in 'single quotes'
* For file names or other things like variables, if they are capitalized in the video, make them capital in the caption.

## Status

Captions and videos go through a number of statuses in their life cycle and will be tracked below.

* **Creation** - This step is performed in YouTube's caption tool, in order to modify their auto-generated captions, fixing incorrect words, timing, and spacing.
* **Final** - This is the final step for the caption files, meant to ensure that everything is correct. Must be done after all have been finished.
* **Complete** - 🎉

| Video | Language | Creation | Final | Complete |
|:-----:|:--------:|:--------:|:-----:|----------|
| 01    | Shared   | √        | √     |          |
| 02    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 03    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 04    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 05    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 06    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 07    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 08    |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 09-10 |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
| 11-12 |          |          |       |          |
|       | Ruby     | √        | √     |          |
|       | Java     | √        | √     |          |
