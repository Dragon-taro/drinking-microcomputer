import { combineReducers } from "redux";
import { partyReducer } from "./party/reducer";

const rootReducer = combineReducers({
  party: partyReducer
});

export default rootReducer;
