import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';

  const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'To ความเชี่ยวชาญเฉพาะทางของแพทย์' };
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Backstage'}`}
        subtitle="ท่านสามารถบันทึกข้อมูลของท่านได้ที่นี้"
      >
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
            <MenuItem 
            onClick={handleClose}
            component={RouterLink}
            to="/"
            >Logout</MenuItem>
          </Menu>
          <p></p>
        </div>
      </Header>
      <Content>
        <ContentHeader title="รายชื่อแพทย์เฉพาะทาง :">
          <Link component={RouterLink} to="/user">
            <Button variant="contained" color="primary">
              บันทึกความเชี่ยวชาญ
           </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
      </Content>
    </Page>
  );
};

export default WelcomePage;
