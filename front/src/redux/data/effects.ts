import { Dispatch } from "redux";
import { get } from "../utils/api";
import { setData, loadingData, setSingleData } from "./action";
import { Data } from "../../entity/data";
import * as io from "socket.io-client";

export const getData = () => async (dispatch: Dispatch) => {
  dispatch(loadingData());
  const data = await get<Data[]>("data");
  dispatch(setData(data));
};

export const subscribeData = () => (dispatch: Dispatch) => {
  const socket = io.connect("https://drinking.dragon-taro.dev/ws");

  socket.on("message", (data: Data) => {
    dispatch(setSingleData(data));
  });
};
