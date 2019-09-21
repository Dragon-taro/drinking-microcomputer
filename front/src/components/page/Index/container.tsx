import { connect } from "react-redux";
import { finishParty } from "../../../redux/party/effects";
import NewParty from "./presentation";
import { DispatchProps } from "./type";

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  finishParty: () => {
    dispatch(finishParty());
  }
});

export default connect(
  null,
  mapDispatchToProps
)(NewParty);
