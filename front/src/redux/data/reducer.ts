import { SET_DATA, LODING_DATA, SET_SINGLE_DATA } from "./constant";
import { DataList } from "../../entity/data";

const init: DataList = {
  data: [],
  isLoading: true
};

export const dataReducer = (state: DataList = init, action: any): DataList => {
  const { type, payload } = action;

  switch (type) {
    case SET_DATA:
      return {
        data: payload,
        isLoading: false
      };
    case SET_SINGLE_DATA:
      return {
        data: [payload, ...state.data],
        isLoading: false
      };
    case LODING_DATA:
      return {
        ...state,
        isLoading: true
      };
    default:
      return state;
  }
};
