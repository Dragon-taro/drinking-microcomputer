import * as React from "react";
import { Props } from "./type";

const NewParty: React.FunctionComponent<Props> = ({ postParty }) => {
  return (
    <div>
      <p>まだ飲み会がありません。下のボタンから飲み会を作成してください。</p>
      <div>
        <button onClick={postParty}>飲み会を作成する</button>
      </div>
    </div>
  );
};

export default NewParty;
