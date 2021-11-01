import moment from 'moment';
import React, { useEffect } from 'react';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Divider,
  Typography
} from '@mui/material';
import QRCode from 'easyqrcodejs';
import { codeoptions } from '../../utils/codestate';
import { useRecoilValue } from 'recoil';

const CodeGeneratorResult = (props) => {
  const codeNode = React.createRef<HTMLDivElement>();
  const options = useRecoilValue(codeoptions);

  useEffect(() => {
    console.log('useffect REsult');
    codeNode.current.innerHTML = '';
    const qrCode = new QRCode(codeNode.current, options);
  }, [options]);

  return (
    <Card {...props}>
      <CardContent>
        <Box
          sx={{
            alignItems: 'center',
            display: 'flex',
            flexDirection: 'column'
          }}
        >
          <div ref={codeNode}></div>
        </Box>
      </CardContent>
      <Divider />
      <CardActions>
        download
        {/*  <Button color="primary" fullWidth variant="text">
        Upload picture
      </Button>
 */}
      </CardActions>
    </Card>
  );
};

export default CodeGeneratorResult;
