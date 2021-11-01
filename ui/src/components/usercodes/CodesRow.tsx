import { useEffect, createRef } from 'react';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import IconButton from '@mui/material/IconButton';
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import {
  Avatar,
  Box,
  Checkbox,
  TableCell,
  TableRow,
  Typography
} from '@mui/material';
import moment from 'moment';
import QRCode from 'easyqrcodejs';

const CodesRow = ({
  code,
  selectedcodeIds,
  handleDelete,
  handleDashboardView,
  handleEditView,
  handleSelectOne,
  ...props
}) => {
  const codeNode = createRef<HTMLDivElement>();

  useEffect(() => {
    if (codeNode != null) {
      const qr = new QRCode(codeNode.current, code.code.options);
      qr.resize(26, 26);
      <div ref={codeNode}></div>;
    }
    return () => {};
  }, []);

  return (
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
          <Avatar sx={{ mr: 2 }}>
            <div ref={codeNode}></div>
          </Avatar>
          <Typography color="textPrimary" variant="body1">
            {window.location.origin}/{code.code.key}
          </Typography>
        </Box>
      </TableCell>
      <TableCell>{code.code.value}</TableCell>
      <TableCell>
        {moment.unix(code.created).format('DD/MM/YYYY - hh:mm')}
      </TableCell>
      <TableCell>
        {moment.unix(code.updated).format('DD/MM/YYYY - hh:mm')}
      </TableCell>
      <TableCell>
        <IconButton onClick={() => handleDashboardView(code)}>
          <DashboardIcon></DashboardIcon>
        </IconButton>
        <IconButton onClick={() => handleEditView(code)}>
          <ModeEditIcon></ModeEditIcon>
        </IconButton>
        <IconButton onClick={() => handleDelete(code)}>
          <DeleteOutlineIcon></DeleteOutlineIcon>
        </IconButton>
      </TableCell>
    </TableRow>
  );
};

export default CodesRow;
