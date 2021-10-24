/* eslint-disable */

import { useState } from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import PerfectScrollbar from 'react-perfect-scrollbar';
import {
  Avatar,
  Box,
  Card,
  Checkbox,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TablePagination,
  TableRow,
  Typography
} from '@mui/material';
import getInitials from '../../utils/getInitials';

const CodesListResults = ({ codes, ...rest }) => {
  console.log(codes);
  const [selectedcodeIds, setSelectedcodeIds] = useState([]);
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(0);

  const handleSelectAll = (event) => {
    let newSelectedcodeIds;

    if (event.target.checked) {
      newSelectedcodeIds = codes.map((code) => code.id);
    } else {
      newSelectedcodeIds = [];
    }

    setSelectedcodeIds(newSelectedcodeIds);
  };

  const handleSelectOne = (event, id) => {
    const selectedIndex = selectedcodeIds.indexOf(id);
    let newSelectedcodeIds = [];

    if (selectedIndex === -1) {
      newSelectedcodeIds = newSelectedcodeIds.concat(selectedcodeIds, id);
    } else if (selectedIndex === 0) {
      newSelectedcodeIds = newSelectedcodeIds.concat(selectedcodeIds.slice(1));
    } else if (selectedIndex === selectedcodeIds.length - 1) {
      newSelectedcodeIds = newSelectedcodeIds.concat(
        selectedcodeIds.slice(0, -1)
      );
    } else if (selectedIndex > 0) {
      newSelectedcodeIds = newSelectedcodeIds.concat(
        selectedcodeIds.slice(0, selectedIndex),
        selectedcodeIds.slice(selectedIndex + 1)
      );
    }

    setSelectedcodeIds(newSelectedcodeIds);
  };

  const handleLimitChange = (event) => {
    setLimit(event.target.value);
  };

  const handlePageChange = (event, newPage) => {
    setPage(newPage);
  };

  return (
    <Card {...rest}>
      <PerfectScrollbar>
        <Box sx={{ minWidth: 1050 }}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell padding="checkbox">
                  <Checkbox
                    checked={selectedcodeIds.length === codes.length}
                    color="primary"
                    indeterminate={
                      selectedcodeIds.length > 0 &&
                      selectedcodeIds.length < codes.length
                    }
                    onChange={handleSelectAll}
                  />
                </TableCell>
                <TableCell>url</TableCell>
                <TableCell>clicked</TableCell>
                <TableCell>link</TableCell>
                <TableCell>created</TableCell>
                <TableCell>updated</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {codes.slice(0, limit).map((code) => (
                <TableRow
                  hover
                  key={code.id}
                  selected={selectedcodeIds.indexOf(code.id) !== -1}
                >
                  <TableCell padding="checkbox">
                    <Checkbox
                      checked={selectedcodeIds.indexOf(code.id) !== -1}
                      onChange={(event) => handleSelectOne(event, code.id)}
                      value="true"
                    />
                  </TableCell>
                  <TableCell>
                    <Box
                      sx={{
                        alignItems: 'center',
                        display: 'flex'
                      }}
                    >
                      <Avatar src={code.avatarUrl} sx={{ mr: 2 }}>
                        {getInitials(code.name)}
                      </Avatar>
                      <Typography color="textPrimary" variant="body1">
                        http://localhost:8080/{code.code.key}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>{code.user}</TableCell>
                  <TableCell>{code.code.value}</TableCell>
                  <TableCell>
                    {moment.unix(code.created).format('DD/MM/YYYY - hh:mm')}
                  </TableCell>
                  <TableCell>
                    {moment.unix(code.updated).format('DD/MM/YYYY - hh:mm')}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Box>
      </PerfectScrollbar>
      <TablePagination
        component="div"
        count={codes.length}
        onPageChange={handlePageChange}
        onRowsPerPageChange={handleLimitChange}
        page={page}
        rowsPerPage={limit}
        rowsPerPageOptions={[5, 10, 25]}
      />
    </Card>
  );
};

CodesListResults.propTypes = {
  codes: PropTypes.array.isRequired
};

export default CodesListResults;
