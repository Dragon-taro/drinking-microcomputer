import { combineReducers } from "redux";
import { partyReducer } from "./party/reducer";
import { dataReducer } from "./data/reducer";

const rootReducer = combineReducers({
  party: partyReducer,
  data: dataReducer
});

export default rootReducer;
