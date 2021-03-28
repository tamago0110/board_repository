import React, { useState } from "react";
import styles from "./EditProfile.module.css";
import Modal from "react-modal";
import { useSelector, useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import AddAPhotoIcon from "@material-ui/icons/AddAPhoto";
import DoneOutlineIcon from "@material-ui/icons/DoneOutline";
import CancelIcon from "@material-ui/icons/Cancel";
import UpdateIcon from "@material-ui/icons/Update";
import { GrEdit } from "react-icons/gr";
import { MdFiberNew } from "react-icons/md";

import {
  fetchAsyncUpdateProfs,
  resetIsOpenPutProfile,
  selectIsOpenPutProfile,
} from "./userSlice";

import { PUT_PROFILE, PROFILE } from "../types";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    doneIcon: {
      marginTop: theme.spacing(1.3),
      marginBottom: theme.spacing(1.3),
      cursor: "none",
    },
    grow: {
      flexGrow: 1,
    },
    name: {
      color: "#3f51b5",
    },
  })
);

const customStyles = {
  content: {
    top: "55%",
    left: "50%",

    width: 300,
    height: 300,

    transform: "translate(-50%, -50%)",
  },
};

const EditProfile: React.FC<PROFILE> = ({ id, user_id, name, image }) => {
  const classes = useStyles();
  const dispatch: AppDispatch = useDispatch();
  const isOpenPutProfile = useSelector(selectIsOpenPutProfile);

  const [editName, setEditName] = useState("");
  const [img, setImg] = useState<File | null>(null);

  const updateProfile = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();
    const packet: PUT_PROFILE = {
      id: id,
      putProfile: { name: editName, image: img },
    };
    dispatch(fetchAsyncUpdateProfs(packet));
    setEditName("");
    setImg(null);
    dispatch(resetIsOpenPutProfile());
  };

  const handleEditPicture = () => {
    const fileInput = document.getElementById("imageInput");
    fileInput?.click();
  };

  return (
    <>
      <Modal
        isOpen={isOpenPutProfile}
        onRequestClose={() => {
          setEditName("");
          setImg(null);
          dispatch(resetIsOpenPutProfile());
        }}
        style={customStyles}
      >
        <p className={styles.nameArea}>
          <GrEdit />
          Old name : <br />
          <a className={classes.name}>{name}</a>
        </p>
        <p className={styles.nameArea}>
          <MdFiberNew />
          New name : <br />
          <a className={classes.name}>{editName}</a>
        </p>
        <br />
        <TextField
          InputLabelProps={{
            shrink: true,
          }}
          required
          fullWidth
          label="New Name"
          type="text"
          value={editName}
          onChange={(e) => setEditName(e.target.value)}
        />
        <input
          type="file"
          id="imageInput"
          hidden={true}
          onChange={(e) => setImg(e.target.files![0])}
        />
        <br />
        <div className={styles.imageArea}>
          {img ? (
            <DoneOutlineIcon color="disabled" className={classes.doneIcon} />
          ) : (
            <IconButton onClick={handleEditPicture}>
              <AddAPhotoIcon color="secondary" />
            </IconButton>
          )}
        </div>
        <Toolbar>
          <Button
            variant="contained"
            color="secondary"
            onClick={() => {
              setEditName("");
              setImg(null);
              dispatch(resetIsOpenPutProfile());
            }}
          >
            <CancelIcon />
            Cancel
          </Button>
          <div className={classes.grow} />
          <Button
            disabled={!editName}
            variant="contained"
            color="secondary"
            type="submit"
            onClick={updateProfile}
          >
            <UpdateIcon />
            Update
          </Button>
        </Toolbar>
      </Modal>
    </>
  );
};

export default EditProfile;
