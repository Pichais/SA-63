import React, { useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import Button from '@material-ui/core/Button';

import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { EntOrgan } from '../../api/models/EntOrgan';
import { EntTypeDisease } from '../../api/models/EntTypeDisease';
import { Select } from '@material-ui/core';
import { EntPhysician } from '../../api/models/EntPhysician';
import { ContentHeader } from '@backstage/core';



const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);


export default function MenuAppBar() {
  const classes = useStyles();
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [physicians, setphysicians] = useState<EntPhysician[]>([]);
  const [organs, setOrgans] = React.useState<EntOrgan[]>([]);
  const [typeDiseases, setTypeDiseases] = React.useState<EntTypeDisease[]>([]);


  const [physicianid, setphysicianid] = useState(Number);
  const [organid, setorganid] = useState(Number);
  const [typediseaseid, settypediseaseid] = useState(Number);





  // Lifecycle Hooks
  useEffect(() => {

    const getOrgans = async () => {
      const or = await api.listOrgan({ limit: 10, offset: 0 });
      setOrgans(or);
    };
    getOrgans();

    const getTypeDiseases = async () => {
      const ty = await api.listTypedisease({ limit: 10, offset: 0 });
      setTypeDiseases(ty);
    };
    getTypeDiseases();

    const getPhysicians = async () => {
      const or = await api.listPhysician({ limit: 10, offset: 0 });
      setphysicians(or);
    };
    getPhysicians();

  }, [loading]);


  const CreateSpacialitie = async () => {
    const spacialitie = {
      organdata: organid,
      physiciandata: physicianid,
      typediseasedata: typediseaseid,
    }
    console.log(spacialitie);
    const res: any = await api.createSpaciality({ spaciality: spacialitie });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const physicianid_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setphysicianid(event.target.value as number);
  };

  const organ_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setorganid(event.target.value as number);
  };

  const type_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    settypediseaseid(event.target.value as number);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h5" className={classes.title}>
            ข้อมูลความเชี่ยวชาญเฉพาะทาง
          </Typography>
          {auth && (
            <div>
              <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={open}
                onClose={handleClose}
              >
                <MenuItem onClick={handleClose}>Logout</MenuItem>
              </Menu>
            </div>
          )}
        </Toolbar>
      </AppBar>

      <ContentHeader title="">
        {status ? (
          <div>
            {alert ? (
              <Alert severity="success">
                บันทึกข้อมูลสำเร็จ!
              </Alert>
            ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  บันทึกข้อมูลล้มเหลว!
                </Alert>
              )}
          </div>
        ) : null}
      </ContentHeader>

      <AppBar position="static" color='inherit' background-color="inherit">
        <Grid container alignItems="center" spacing={2} >
          <Grid item xs={12}></Grid>
          <Grid item xs={12}></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>NAME</h3></Grid>
          <Grid item xs={2}>
            <Select
              labelId="user_id-label"
              label="User"
              id="user_id"
              value={physicianid}
              onChange={physicianid_id_handleChange}
              style={{ width: 200 }}
            >
              {physicians.map((item: EntPhysician) =>
                <MenuItem value={item.id}>{item.physicianName}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>E-mail</h3></Grid>
          <Grid item xs={2}>
            <Select
              labelId="user_id-label"
              label="User"
              id="user_id"
              value={physicianid}
              onChange={physicianid_id_handleChange}
              style={{ width: 200 }}
            >
              {physicians.map((item: EntPhysician) =>
                <MenuItem value={item.id}>{item.physicianEmail}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>โปรดเลือกความเชี่ยวชาญด้านอวัยวะ</h3></Grid>
          <Grid item xs={2}>
            <Select
              labelId="user_id-label"
              label="Organ"
              id="Organ_id"
              value={organid}
              onChange={organ_id_handleChange}
              style={{ width: 200 }}
            >
              {organs.map((item: EntOrgan) =>
                <MenuItem value={item.id}>{item.organName}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3>โปรดเลือกชนิดของโรค</h3></Grid>
          <Grid item xs={2}>
            <Select
              labelId="Typedisease_id-label"
              label="Typedisease"
              id="Typedisease_id"
              value={typediseaseid}
              onChange={type_id_handleChange}
              style={{ width: 200 }}
            >
              {typeDiseases.map((item: EntTypeDisease) =>
                <MenuItem value={item.id}>{item.diseaseName}</MenuItem>)}
            </Select>
          </Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><h3></h3></Grid>
          <Grid item xs={2}></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={3}><p></p></Grid>
          <Grid item xs={3}></Grid>
          <Grid item xs={3}></Grid>
          <Grid item xs={3}><p></p></Grid>

          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}>
            <Button
              onClick={() => {
                CreateSpacialitie();
              }}

              color="inherit"
              style={{ width: 200 }}
              variant="contained">

              SAVE
            </Button>
          </Grid>
          <Grid item xs={1}></Grid>
          <Grid item xs={1}>
            <Button
              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/login"
              variant="contained"
            >
              Back
            </Button>
          </Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>
          <Grid item xs={1}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>
          <Grid item xs={2}><p></p></Grid>

        </Grid>
      </AppBar>


      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
          </Typography>
        </Toolbar>
      </AppBar>
    </div>



  );
}