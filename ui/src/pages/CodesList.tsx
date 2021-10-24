import { Helmet } from 'react-helmet';
import { useEffect, useState } from 'react';
import { Box, Container } from '@mui/material';
import CustomerListResults from '../components/usercodes/CodesListResults';
import CustomerListToolbar from '../components/usercodes/CodesListToolbar';
import { CodesApiService } from '../gen/api/services/CodesApiService';

const CodesList = () => {
  const [codes, setCodes] = useState([]);

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
