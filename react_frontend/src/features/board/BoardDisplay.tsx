import React, { useState, useEffect } from "react";
import styles from "./BoardDisplay.module.css";
import { Theme, createStyles, makeStyles } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import DeleteForeverIcon from "@material-ui/icons/DeleteForever";
import SendIcon from "@material-ui/icons/Send";

import { useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store";
import { READ_BOARD, PROFILE } from "../types";
import {
  fetchAsyncPostDisplay,
  fetchAsyncPostLead,
  fetchAsyncGetBoards,
  setIsOpenPostBoard,
} from "./boardSlice";
import { fetchAsyncGetSpecificProf } from "../user/userSlice";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    grow: {
      flexGrow: 1,
    },
    fabButton: {
      position: "absolute",
      top: -10,
      left: 0,
      right: 0,
      margin: "0 auto",
    },
    avatar: {
      width: theme.spacing(3),
      height: theme.spacing(3),
      display: "inline-block",
      textAlign: "center",
    },
  })
);

const BoardDisplay: React.FC<READ_BOARD> = ({
  id,
  created_by,
  title,
  content,
}) => {
  const classes = useStyles();
  const dispatch: AppDispatch = useDispatch();
  const [creatorPro, setCreatorPro] = useState<PROFILE>({
    id: 0,
    user_id: "",
    name: "",
    image: "",
  });

  const getProf = async () => {
    const res = await dispatch(fetchAsyncGetSpecificProf(created_by));
    if (fetchAsyncGetSpecificProf.fulfilled.match(res)) {
      setCreatorPro(res.payload);
    }
  };

  useEffect(() => {
    created_by && getProf();
  }, [id]);

  return (
    <div className={styles.boardContainer}>
      <div className={styles.boardHeader}>
        {title === "" ? (
          <h2>No boards</h2>
        ) : (
          <h2 className={styles.boardTitle}>{title}</h2>
        )}
        {created_by && (
          <p className={styles.createdBy}>
            created by{" "}
            <Avatar src={creatorPro.image} className={classes.avatar} />
          </p>
        )}
      </div>
      <div className={styles.boardContent}>
        <h4>{content}</h4>
      </div>
      <div className={styles.buttonContainer}>
        <Toolbar>
          {id === 0 ? (
            <a />
          ) : (
            <IconButton
              edge="start"
              color="inherit"
              aria-label="open drawer"
              onClick={async () => {
                await dispatch(fetchAsyncPostDisplay({ board_id: id }));
                await dispatch(fetchAsyncGetBoards());
              }}
            >
              <DeleteForeverIcon />
            </IconButton>
          )}
          <Fab
            color="secondary"
            aria-label="add"
            className={classes.fabButton}
            onClick={() => {
              dispatch(setIsOpenPostBoard());
            }}
          >
            <AddIcon />
          </Fab>
          <div className={classes.grow} />
          {id === 0 ? (
            <a />
          ) : (
            <IconButton
              edge="end"
              color="inherit"
              onClick={async () => {
                await dispatch(fetchAsyncPostLead({ producer: created_by }));
                await dispatch(fetchAsyncGetBoards());
              }}
            >
              <SendIcon />
            </IconButton>
          )}
        </Toolbar>
      </div>
    </div>
  );
};

export default BoardDisplay;
