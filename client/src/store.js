import { createStore, applyMiddleware, compose } from 'redux'
import rootReducer from "./reducers"

function logger({ getState }) {
  return next => action => {
    console.log('will dispatch', action)

    // Call the next dispatch method in the middleware chain.
    const returnValue = next(action)

    // console.log('state after dispatch', getState())

    // This will likely be the action itself, unless
    // a middleware further in chain changed it.
    return returnValue
  }
}

const store = createStore(
  rootReducer, {},
  compose(applyMiddleware(logger), window.devToolsExtension ? window.devToolsExtension() : f => f)
)

export default store