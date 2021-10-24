import { Helmet } from 'react-helmet';
import { useEffect } from 'react';
import { Box, Container } from '@mui/material';
import CustomerListResults from '../components/usercodes/CodesListResults';
import CustomerListToolbar from '../components/usercodes/CodesListToolbar';
import { CodesApiService } from '../gen/api/services/CodesApiService.ts';

const CodesList = () => {
  let codes = [];

  useEffect(() => {
    CodesApiService.getCodesFm().then((response) => {
      console.log(response);
      codes = response;
    });
    /*     return () => {
    };
 */
  }, []);

  return (
    <>
      <Helmet>
        <title>Customers | Material Kit</title>
      </Helmet>
      <Box
        sx={{
          backgroundColor: 'background.default',
          minHeight: '100%',
          py: 3
        }}
      >
        <Container maxWidth={false}>
          <CustomerListToolbar />
          <Box sx={{ pt: 3 }}>
            <CustomerListResults codes={codes} />
          </Box>
        </Container>
      </Box>
    </>
  );
};

export default CodesList;
