import { connect } from "react-redux";
import { postParty } from "../../../redux/party/effects";
import NewParty from "./presentation";
import { DispatchProps } from "./type";

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  postParty: () => {
    dispatch(postParty());
  }
});

export default connect(
  null,
  mapDispatchToProps
)(NewParty);
