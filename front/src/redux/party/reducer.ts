import { Party } from "../../entity/party";
import { SET_PARTY, LOADING_PARTY } from "./constant";

const init: Party = {
  isLoading: true
};

export const partyReducer = (state: Party = init, action: any): Party => {
  const { type, payload } = action;

  switch (type) {
    case SET_PARTY:
      return {
        ...payload,
        isLoading: false
      };
    case LOADING_PARTY:
      return {
        ...state,
        isLoading: true
      };
    default:
      return state;
  }
};
