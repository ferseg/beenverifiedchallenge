# Beenverified Challenge

## Summary
Basic API with golang.

## Prerequisites
1. The following programs must be installed before you run the API
    * [Golang](https://golang.org/dl/) (Version 1.6.x)
    * [Glide](https://github.com/Masterminds/glide)
2. The *GOPATH* variable must be set


## Instalation
Before compiling the program you must download the repo in your *GOPATH/src* location

### Linux and MacOSX
1. Access to your source folder in the GOPATH: `cd $GOPATH/src`
2. Clone the repository: `git clone git@github.com:ferseg/beenverifiedchallenge.git` or `git clone https://github.com/ferseg/beenverifiedchallenge.git`
3. Go inside the downloaded folder: `cd beenverifiedchallenge`
4. Execute the glide update: `glide update`
5. Install the program: `go install`

## Running the program

### Linux and MacOSX
To run the program use the following command: `$GOPATH/bin/beenverifiedchallenge`
* The server starts listening
* Open the browser

## API
### Get songs infromation by artist, song or genre
Returns the information of all songs that match with a criteria.
URL: `localhost:8000/song/[criteria]` replace the *criteria* with your search text.
### Get songs information which has a length in a range
Returns the information of all songs that has a length in a range.
URL: `localhost:8000/song/[min]/[max]` the *min* value means the minimum length of the song and the *max* value means the maximum length of the song
### Get genres information
Returns the information of all genres, the information includes:
* name The name of the genre
* songQuantity The amount of songs that the genre has.
* totalTime The sum of the lenghts of all songs with that genre.
URL: `localhost:8000/genreInfo`