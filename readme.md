# Vid_Trimmer

This is a service written in Go to save me from downloading a full video when I only need a part from the video.

I'm a low budget and lazy developer, so I don't have the money or time to make it a tool for everyone. If you have the 
time and cash to create a service from it, please do.

The project uses [github.com/jtguibas/cinema](https://github.com/jtguibas/cinema) to trim videos.

## Supported video sites
- Youtube videos using [github.com/kkdai/youtube](https://github.com/kkdai/youtube)
- twitter videos using [github.com/dghubble/go-twitter](https://github.com/dghubble/go-twitter)
- video url 

## Generate Download link

This endpoint generates a download link for the trimmed video.

Path: [GET]`/download`

### Request Parameters 
- **url**: the url to download from
- **format**: the format to return the file
- **start**: the second `s` to start the trim
- **end**: the second `e` to end the trim

### Response

This returns a json object that contains `link` and `state` fields.
 
`link` is the download link of the trimmed video.

`state` gives info about the state of the video.

- **NULL**: no file yet
- **PENDING**: trimming started, but not complete
- **DONE**: video trimmed successfully

The frontend will have to make a request periodically till the `state` is `DONE`. This is because I don't have time to add 
websocket (remember, I'm a lazy dev)


```json
{
"link":"localhost:8080/download/41c6d76521.gif",
"state":"DONE"
}
```

## Download File

This endpoint downloads the trimmed video. This is generated by, `Generate Download link`

Path: [GET]`/download/{key}`