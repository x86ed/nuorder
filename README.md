# NuOrder Interview Homework 
#### prepared by [Adam Siegel](mailto:adam.siegel.is@gmail.com)

## Prerequisites

* **Golang** installed (I built it using `go1.13.9 darwin/amd64` but any `>1.0` release should work)
* **Node.js** installed
* **npx** installed (you may not need this to run but I built of the latest version of `create-react-app`)

## To Run the App

1. checkout the repo locally
1. navigate to the project root from your shell
1. Run 
```bash
$> npm start
```

the browser should open to [http://localhost:3000](http://localhost:3000) (if not click that link)

## Hotkeys

#### Enter 
sets the user input to the active selection in the suggestion list

#### UP
moves makes the previous selection in the list the active selection

#### Down
moves makes the next selection in the list the active selection

## Things I'd like to improve
##### (If given more time to do this assignment)

1. **Tests** there aren't enough tests here. I'm not a fan of BDD/TDD but acceptance tests are a good practice for any new feature or component. Also it would be good to integrate the Go and Node runners.
1. **Go Proxy** this should have better error handling & logging. The implementation is rough and completes the assignment but I'd never use it as a production service.
1. **API Integration with GitHub** this should be authenticated instead of anonymous. (too large scope for this project)
1. **AJAX** the API requests should probably cached server side intitally & be refreshed when the input to search changes.
1. **Flux/Redux** this project should have an event driven in browser datastore to follow best React practices

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode and starts the go proxy app.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.
Open [http://localhost:3000/hub](http://localhost:3000/hub) to view the github API proxy

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.
