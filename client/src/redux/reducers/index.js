import { ADD_CURR_PROJECT } from "../constants/action-types";
import { SET_CURR_PROJECT } from "../constants/action-types";
import { SET_CURR_POPUP_ID } from "../constants/action-types";
import { SAVE_PROJECT_TO_KV } from "../constants/action-types";
import { LOAD_PROJECT_FROM_KV } from "../constants/action-types";
import { RETURN_LOAD_PROJECT } from "../constants/action-types";
import { GENERATE_CURR_PROJECT } from "../constants/action-types"
import { DELIVER_GENERATED_TAR } from "../constants/action-types"
import { socket } from '../../index';


const initialState = {
  currPopupID: null,
  projects: [],
  currProject: null,
};
function rootReducer(state = initialState, action) {
  //Add the current project to the array of saved projects
  if (action.type === ADD_CURR_PROJECT) {
    return Object.assign({}, state, {
      projects: state.projects.concat(action.payload)
    });
  }
  //Makes the project the current project in redux
  else if (action.type === SET_CURR_PROJECT) {
    return Object.assign({}, state, {
      currProject: action.payload
    });
  }
  //Set the popup id of the clicked plugin in redux for use
  //as a global
  else if (action.type === SET_CURR_POPUP_ID) {
    return Object.assign({}, state, {
      currPopupID: action.payload
    });
  }
  //Retreives the generated tar from the backend server
  else if (action.type === DELIVER_GENERATED_TAR) {
    return Object.assign({}, state, {
      tar: action.payload
    });
  }
  //Retreives the loaded project from the backend server
  else if (action.type === RETURN_LOAD_PROJECT) {
    return Object.assign({}, state, {
      projects: state.projects.concat(action.payload)
    })
  }
  //Emits the server to call GENERATE_PROJECT
  else if (action.type === GENERATE_CURR_PROJECT) {
    socket && socket.emit('GENERATE_PROJECT', action.payload);
  }
  //Emits the server to call SEND_SAVE_PROJECT
  else if (action.type === SAVE_PROJECT_TO_KV) {
    socket && socket.emit('SEND_SAVE_PROJECT', action.payload)
  }
  //Emits the server to call SEND_LOAD_PROJECT
  else if (action.type === LOAD_PROJECT_FROM_KV) {
    socket && socket.emit('SEND_LOAD_PROJECT', action.payload)
  }
  return state;
}
export default rootReducer;