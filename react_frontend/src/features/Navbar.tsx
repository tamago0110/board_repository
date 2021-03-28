import React from "react";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import Avatar from "@material-ui/core/Avatar";
import { AppDispatch } from "../app/store";
import { useSelector, useDispatch } from "react-redux";
import { selectLoginUser, setIsOpenPutProfile } from "./user/userSlice";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    title: {
      flexGrow: 1,
      marginLeft: theme.spacing(5),
    },
    signOut: {
      marginRight: theme.spacing(2),
      backgroundColor: "transparent",
      border: "none",
      outline: "none",
      color: "white",
      cursor: "pointer",
    },
    avatar: {
      width: theme.spacing(3),
      height: theme.spacing(3),
      marginRight: theme.spacing(1),
      cursor: "pointer",
    },
  })
);

const Navbar: React.FC = () => {
  const classes = useStyles();
  const dispatch: AppDispatch = useDispatch();
  const loginUser = useSelector(selectLoginUser);

  const Logout = () => (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    localStorage.removeItem("localJWT");
    window.location.href = "/";
  };

  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          {loginUser.name === "" ? (
            <Typography variant="h5" className={classes.title}>
              Hello, Guest!!
            </Typography>
          ) : (
            <Typography variant="h5" className={classes.title}>
              Hello, {loginUser.name}!!
            </Typography>
          )}

          <Avatar
            alt="who?"
            src={loginUser.image}
            className={classes.avatar}
            onClick={() => dispatch(setIsOpenPutProfile())}
          />

          <button className={classes.signOut} onClick={Logout()}>
            <ExitToAppIcon />
          </button>
        </Toolbar>
      </AppBar>
    </div>
  );
};

export default Navbar;
