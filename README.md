## Intro

A steganography project just for fun. Inspired by the show Mr. Robot, in which Eliott hides secret documents within audio CDs.  

This program lets you do the same thing, but you get to hide "secret" documents inside audio files.  

The source document/files are first encrypted, and then injected into the audio file.   

Note that you don't *have* to use an audio file to hide your data within. Any file will work, as long as its large enough to hold your secret data!  

&nbsp;

## Building / installing
### Linux/MacOS: 
Check out the repo and run `./build.sh`  

&nbsp;

## Usage examples
### Hiding a document inside an audio file
`wavhide -gap 512 -audioFile example.mp3 -secretFile secret.docx`

At this point you are free to delete your original document (`secretFile`). All the data is now stored/hidden within `example-wavhidden.mp3`!  

&nbsp;


### Retrieving a document hidden inside an audio file
`wavreveal -gap 512 -audioFile example-wavhidden.mp3 -outputFile my-file-revealed.docx`

&nbsp;


## Tips / notes
- Smaller "secret" files = less space needed = less chance of audible artifacts.  Maybe compress your files first, or use plain text files for best results.
- File size output will be identical to the original/source song.  


## Gap size
The `-gap` parameter defines how much space to leave intouched between bytes injected into the target audio file.

It's a trade off. A smaller gap size leads to more artifacts in the song, but gives you more space to hide data. A longer gap reduces the number of artifacts in the song, but gives you less space for hiding data.

Remember the gap you use to `wavhide` the secret file! You must use the same gap size when revealing. 

Best practice us to use the largest possible gap, since this spreads out the hidden data as sparsely as possible, which leads to the minimum amount of detectable artifacts in the resulting file. 


## Other best practices
For your `-audioFile`, use a file of your creation. If you use a publically-available file, an attacker could run a diff against the public files and your files to discover the files' gap length. They still couldn't extract the contents though since the source file is encrypted before being hidden in the destination audio file.

.wav files are recommended since they tolerate random bytes without creating as many noticable audio artifact compred to mp3 files. That said, there is *just* enough smarts to avoid destroying the MP3 headers if you want to hide stuff within .mp3 files. 
