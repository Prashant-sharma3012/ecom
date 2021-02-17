import { combineReducers } from 'redux'
import ItemReducer from './item'

export default combineReducers({
  user: ItemReducer
})