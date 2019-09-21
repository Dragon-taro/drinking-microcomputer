import { connect } from "react-redux";
import { finishParty } from "../../../redux/party/effects";
import NewParty from "./presentation";
import { DispatchProps } from "./type";
import { getData } from "../../../redux/data/effects";
import { RootState } from "../../../entity/state";
import { StateProps } from "./type";

const mapStateToProps = (state: RootState): StateProps => ({
  data: state.data
});

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
  finishParty: () => {
    dispatch(finishParty());
  },
  getData: () => {
    dispatch(getData());
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewParty);
