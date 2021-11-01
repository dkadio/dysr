/* eslint-disable */

import { useState, useEffect, createRef } from 'react';
import PropTypes from 'prop-types';
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
import CodesRow from './CodesRow';
import { useNavigate } from 'react-router-dom';
import { CodesApiService } from '../../gen/api/services/CodesApiService';
import { useRecoilState } from 'recoil';
import { codeslist } from '../../utils/codestate';
import { SettingsRemoteRounded } from '@material-ui/icons';

const CodesListResults = ({ codes, ...rest }) => {
  console.log(codes);
  const [selectedcodeIds, setSelectedcodeIds] = useState([]);
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(0);
  const nav = useNavigate();
  const [codesList, setCodes] = useRecoilState(codeslist);

  const handleSelectAll = (event) => {
    let newSelectedcodeIds;

    if (event.target.checked) {
      newSelectedcodeIds = codes.map((code) => code.id);
    } else {
      newSelectedcodeIds = [];
    }
    setSelectedcodeIds(newSelectedcodeIds);
  };

  const handleDashboardView = (item) => {
    nav('/app/dashboard/' + item.id);
  };
  const handleEditView = (item) => {
    nav('/app/generator/' + item.id);
  };
  const handleDelete = (item) => {
    CodesApiService.deleteCodeFm(item.id)
      .then((response) => {
        const newCodes = codesList.filter((code) => code.id !== item.id);
        setCodes(newCodes);
      })
      .catch((err) => {
        console.log(err);
      });
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
                <TableCell>address</TableCell>
                <TableCell>created</TableCell>
                <TableCell>updated</TableCell>
                <TableCell>actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {codes.slice(0, limit).map((code) => (
                <CodesRow
                  key={code.id}
                  code={code}
                  handleDelete={handleDelete}
                  handleDashboardView={handleDashboardView}
                  handleEditView={handleEditView}
                  selectedcodeIds={selectedcodeIds}
                  handleSelectOne={handleSelectOne}
                />
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
