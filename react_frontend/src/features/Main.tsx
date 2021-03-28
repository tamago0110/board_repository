import React, { useEffect } from "react";
import styles from "./Main.module.css";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import NoteIcon from "@material-ui/icons/Note";
import ListIcon from "@material-ui/icons/List";
import SupervisorAccountIcon from "@material-ui/icons/SupervisorAccount";

import { useSelector, useDispatch } from "react-redux";
import {
  selectLoginUser,
  selectProfiles,
  fetchAsyncGetMyProf,
  fetchAsyncGetProfs,
} from "./user/userSlice";
import {
  selectDisplayBoards,
  selectWhichBoard,
  selectMyBoards,
  fetchAsyncGetBoards,
  fetchAsyncGetMyBoards,
} from "./board/boardSlice";
import { AppDispatch } from "../app/store";
import BoardDisplay from "./board/BoardDisplay";
import MyBoard from "./board/MyBoard";
import EditProfile from "./user/EditProfile";
import LeadProfile from "./user/LeadProfile";
import CreateBoard from "./board/CreateBoard";
import EditBoard from "./board/EditBoard";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    grow: {
      flexGrow: 1,
    },
    icons: {
      marginRight: "7px",
      backgroundColor: "transparent",
      border: "none",
      outline: "none",
      fontSize: "20px",
    },
  })
);

const Main: React.FC = () => {
  const classes = useStyles();

  const dispatch: AppDispatch = useDispatch();
  const profiles = useSelector(selectProfiles);
  const loginUser = useSelector(selectLoginUser);
  const myBoards = useSelector(selectMyBoards);
  const displayBoards = useSelector(selectDisplayBoards);
  const whichBoard = useSelector(selectWhichBoard);

  useEffect(() => {
    const fetchBootLoader = async () => {
      await dispatch(fetchAsyncGetMyProf());
      await dispatch(fetchAsyncGetProfs());
      await dispatch(fetchAsyncGetBoards());
      await dispatch(fetchAsyncGetMyBoards());
    };
    fetchBootLoader();
  }, [dispatch]);

  return (
    <>
      <EditProfile
        id={loginUser.id}
        user_id={loginUser.user_id}
        name={loginUser.name}
        image={loginUser.image}
      />
      <Grid container>
        <Grid item xs={8} lg={8}>
          <h3 className={styles.areaTitle}>
            <NoteIcon className={classes.icons} />
            Board inbox
          </h3>
          <div className={styles.appBoard}>
            {displayBoards.length !== 0 ? (
              <BoardDisplay
                id={displayBoards[whichBoard].id}
                created_by={displayBoards[whichBoard].created_by}
                title={displayBoards[whichBoard].title}
                content={displayBoards[whichBoard].content}
              />
            ) : (
              <BoardDisplay id={0} created_by={""} title={""} content={""} />
            )}
          </div>
        </Grid>
        <Grid item xs={6} lg={4}>
          <h3 className={styles.areaTitle}>
            <SupervisorAccountIcon className={classes.icons} />
            Leads list
          </h3>
          <div className={styles.appLeadsList}>
            {profiles.length === 0 ? (
              <p className={styles.sentence}>No leads</p>
            ) : (
              <ul className={styles.leadsContainer}>
                {profiles.map((profile) => (
                  <LeadProfile
                    key={profile.id}
                    id={profile.id}
                    user_id={profile.user_id}
                    name={profile.name}
                    image={profile.image}
                  />
                ))}
              </ul>
            )}
          </div>
          <h3 className={styles.areaTitle}>
            <ListIcon className={classes.icons} />
            Created boards
          </h3>
          <div className={styles.appBoardsList}>
            {myBoards.length === 0 ? (
              <p className={styles.sentence}>No contents</p>
            ) : (
              <ul>
                {myBoards.map((myBoard) => (
                  <MyBoard
                    key={myBoard.id}
                    id={myBoard.id}
                    created_by={myBoard.created_by}
                    title={myBoard.title}
                    content={myBoard.content}
                  />
                ))}
              </ul>
            )}
          </div>
        </Grid>
      </Grid>
      <CreateBoard />
      <EditBoard />
    </>
  );
};

export default Main;
