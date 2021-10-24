import { Helmet } from 'react-helmet';
import { Box, Container, Grid } from '@mui/material';
import CodeGeneratorResult from '../components/codegenerator/CodeGeneratorResult';
import CodeGeneratorSettings from '../components/codegenerator/CodeGeneratorSettings';

const CodeGeneratorPage = () => (
  <>
    <Helmet>
      <title>Account | Material Kit</title>
    </Helmet>
    <Box
      sx={{
        backgroundColor: 'background.default',
        minHeight: '100%',
        py: 3
      }}
    >
      <Container maxWidth="lg">
        <Grid container spacing={3}>
          <Grid item lg={4} md={6} xs={12}>
            <CodeGeneratorResult />
          </Grid>
          <Grid item lg={8} md={6} xs={12}>
            <CodeGeneratorSettings />
          </Grid>
        </Grid>
      </Container>
    </Box>
  </>
);

export default CodeGeneratorPage;
