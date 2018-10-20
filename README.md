# cucumber-school-captions

Generating captions for Cucumber School videos!

Steps

* Use [autosub](https://github.com/agermanidis/autosub) for the initial generation of VTT files. Clean up will need to be done by hand.
  * Example: `autosub -F vtt 02_your_first_scenario/java/Cucumber\ School\ -\ 01x02\ -\ Your\ First\ Scenario\ \(Java\).mp4`
* Run through video with caption file to clean up the inevitable errors. Misspellings, incorrect words, missing lines, and timing issues may all pop up.
* Use ffmpeg to combine the caption file with the video
* Upload
* Success!
