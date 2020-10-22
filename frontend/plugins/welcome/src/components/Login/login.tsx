import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
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
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
       subtitle="ท่านสามารถบันทึกข้อมูลของท่านได้ที่นี้"
     ></Header>
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