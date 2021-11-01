import { Helmet } from 'react-helmet';
import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Box, Container, Grid } from '@mui/material';
import CodeGeneratorResult from '../components/codegenerator/CodeGeneratorResult';
import CodeGeneratorSettings from '../components/codegenerator/CodeGeneratorSettings';
import { CodesApiService } from '../gen/api/services/CodesApiService';
import { useRecoilState, useResetRecoilState } from 'recoil';
import { codeoptions, codevalue } from '../utils/codestate';
import { Columns } from 'react-feather';

const CodeGeneratorPage = () => {
  const { id } = useParams();
  const [options, setOptions] = useRecoilState(codeoptions);
  const [value, setValue] = useRecoilState(codevalue);
  const resetOptions = useResetRecoilState(codeoptions);
  const nav = useNavigate();

  useEffect(() => {
    if (id !== undefined) {
      console.log('Loading Code with id', id);
      CodesApiService.getCodeFm(id)
        .then((response) => {
          setValue(response.code.value);
          setOptions(response.code.options);
          console.log('Set options for given id', response.code.options);
        })
        .catch((err) => {
          console.log(err);
        });
      /*     return () => {
    };
 */
    } else {
      //create a new one and redirect
      console.log('reset options', options.text);
      CodesApiService.createCodeFm({ options: options })
        .then((response) => {
          console.log('Options', options);
          console.log('newly created qrcode ', response);
          setValue(response.code.value);
          setOptions(response.code.options);
          nav('/app/generator/' + response.id);
        })
        .catch((err) => {
          console.log(err);
        });
    }
    return () => {
      console.log(options.text);
      resetOptions();
    };
  }, []);

  const savecode = (e) => {
    console.log('save code ', id);
    console.log('with options: ', options);
    console.log('with value,', value);
    CodesApiService.updateCodeFm(id, { options: options, value: value });
  };

  return (
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
            <Grid item lg={5} md={6} xs={12}>
              <CodeGeneratorResult />
            </Grid>
            <Grid item lg={7} md={6} xs={12}>
              <CodeGeneratorSettings savecode={savecode} />
            </Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
};

export default CodeGeneratorPage;
