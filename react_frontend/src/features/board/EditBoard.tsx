import React from "react";
import Modal from "react-modal";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button";
import CancelIcon from "@material-ui/icons/Cancel";
import UpdateIcon from "@material-ui/icons/Update";
import { useSelector, useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store";

import {
  selectIsOpenPutBoard,
  selectUpdateBoard,
  resetIsOpenPutBoard,
  editMyBoardTitle,
  editMyBoardContent,
  fetchAsyncPutMyBoard,
} from "./boardSlice";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    grow: {
      flexGrow: 1,
    },
  })
);

const customStyles = {
  content: {
    width: 400,
    height: 390,
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
  },
};

const EditBoard: React.FC = () => {
  const classes = useStyles();
  const dispatch: AppDispatch = useDispatch();
  const isOpenPutBoard = useSelector(selectIsOpenPutBoard);
  const updateBoard = useSelector(selectUpdateBoard);

  return (
    <>
      <Modal
        isOpen={isOpenPutBoard}
        onRequestClose={() => dispatch(resetIsOpenPutBoard())}
        style={customStyles}
      >
        <TextField
          InputLabelProps={{
            shrink: true,
          }}
          required
          fullWidth
          label="Title"
          type="text"
          name="title"
          value={updateBoard.putBoard.title}
          onChange={(e) => dispatch(editMyBoardTitle(e.target.value))}
        />
        <TextField
          InputLabelProps={{
            shrink: true,
          }}
          required
          fullWidth
          multiline
          rows={13}
          label="Content"
          type="text"
          name="content"
          value={updateBoard.putBoard.content}
          onChange={(e) => dispatch(editMyBoardContent(e.target.value))}
        />
        <Toolbar>
          <Button
            variant="contained"
            color="secondary"
            onClick={() => {
              dispatch(resetIsOpenPutBoard());
            }}
          >
            <CancelIcon />
            Cancel
          </Button>
          <div className={classes.grow} />
          <Button
            variant="contained"
            color="secondary"
            onClick={() => {
              dispatch(fetchAsyncPutMyBoard(updateBoard));
              dispatch(resetIsOpenPutBoard());
            }}
          >
            <UpdateIcon />
            Update
          </Button>
        </Toolbar>
      </Modal>
    </>
  );
};

export default EditBoard;
