import { Helmet } from 'react-helmet';
import { useEffect, useState } from 'react';
import { Box, Container } from '@mui/material';
import CustomerListResults from '../components/usercodes/CodesListResults';
import CodesListToolbar from '../components/usercodes/CodesListToolbar';
import { CodesApiService } from '../gen/api/services/CodesApiService';
import { codeslist } from '../utils/codestate';
import { useRecoilState } from 'recoil';

const CodesList = () => {
  const [codes, setCodes] = useRecoilState(codeslist);

  useEffect(() => {
    CodesApiService.getCodesFm()
      .then((response) => {
        setCodes(response);
      })
      .catch((err) => {
        console.log(err);
      });
    /*     return () => {
    };
 */
  }, []);

  return (
    <>
      <Helmet>
        <title>Dynamic QR Codes | Statistics Free OpenSource</title>
      </Helmet>
      <Box
        sx={{
          backgroundColor: 'background.default',
          minHeight: '100%',
          py: 3
        }}
      >
        <Container maxWidth={false}>
          <CodesListToolbar />
          <Box sx={{ pt: 3 }}>
            <CustomerListResults codes={codes} />
          </Box>
        </Container>
      </Box>
    </>
  );
};

export default CodesList;
