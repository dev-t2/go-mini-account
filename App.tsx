import { memo } from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import styled from '@emotion/native';

import { Ball } from './src/components';

const StyledGestureHandlerRootView = styled(GestureHandlerRootView)({
  flex: 1,
});

const App = () => {
  return (
    <StyledGestureHandlerRootView>
      <Ball />
    </StyledGestureHandlerRootView>
  );
};

export default memo(App);
