import React from "react";
import styles from "./LeadProfile.module.css";
import { PROFILE } from "../types";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    avatar: {
      width: theme.spacing(4),
      height: theme.spacing(4),
      top: theme.spacing(-1.5),
      display: "inline-block",
      textAlign: "center",
    },
  })
);

const LeadProfile: React.FC<PROFILE> = ({ id, user_id, name, image }) => {
  const classes = useStyles();
  return (
    <li className={styles.listStyle}>
      <Avatar alt="who?" src={image} className={classes.avatar} />
      {name === "" ? (
        <h3 className={styles.leadName}>No name</h3>
      ) : (
        <h3 className={styles.leadName}>{name}</h3>
      )}
    </li>
  );
};

export default LeadProfile;
