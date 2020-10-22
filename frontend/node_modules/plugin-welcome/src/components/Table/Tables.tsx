import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntSpaciality } from '../../api/models/EntSpaciality';
import { EntOrgan } from '../../api/models/EntOrgan';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [spacialitys, setSpacialitys] = useState<EntSpaciality[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getSpacialitys = async () => {
      const res = await api.listSpaciality({ limit: 10, offset: 0 });
      setLoading(false);
      setSpacialitys(res);
    };
    getSpacialitys();
  }, [loading]);

  const deleteSpacialitys = async (id: number) => {
    const res = await api.deleteSpaciality({ id: id });
    setLoading(true);
  };

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">Organ</TableCell>
            <TableCell align="center">typedisease</TableCell>
            <TableCell align="center">Physician</TableCell>
            
          </TableRow>
        </TableHead>
        <TableBody>
         
          {spacialitys.map((item:any)=> (
              <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.edges?.ogan?.organName}</TableCell>
                <TableCell align="center">{item.edges?.type?.diseaseName}</TableCell>
                <TableCell align="center">{item.edges?.user?.physicianName}</TableCell>
                <TableCell align="center">
                  <Button
                    onClick={() => {
                      deleteSpacialitys(item.id);
                    }}
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}