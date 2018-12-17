imgseq is an image sequence editor

Currently there is only one filter: timeshift. This filter takes the input image sequence and replaces each pixel with a pixel in the sequence as determined by the grayscale value of a filter image sequence. 

Usage example:
```imgseq input01.png timeshift filterimg=./filterimg01.png:range=10```

This takes the image sequence beginning at input01.png and applies the timeshift filter with a filter sequence beginning at filterimg01.png. The range is the maximum distance the shift will replace pixels with.

Output images will be dumped in current directory as out###.png
