import { SET_DATA } from "./constant";
import { Data } from "../../entity/data";

const init = null;

export const dataReducer = (
  state: Data[] | null = init,
  action: any
): Data[] | null => {
  const { type, payload } = action;

  switch (type) {
    case SET_DATA:
      return payload;
    default:
      return state;
  }
};
