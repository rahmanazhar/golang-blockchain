import React, { useState, useEffect, useCallback } from 'react';
import {
  Box,
  Container,
  VStack,
  Heading,
  Text,
  Button,
  Input,
  Card,
  CardBody,
  Badge,
  useToast,
} from '@chakra-ui/react';
import axios from 'axios';

function App() {
  const [blocks, setBlocks] = useState([]);
  const [isValid, setIsValid] = useState(true);
  const [newBlockData, setNewBlockData] = useState('');
  const toast = useToast();

  const fetchBlockchain = useCallback(async () => {
    try {
      const response = await axios.get('http://localhost:8080/blockchain');
      setBlocks(response.data.blocks);
      setIsValid(response.data.valid);
    } catch (error) {
      toast({
        title: 'Error fetching blockchain',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    }
  }, [toast]);

  const addBlock = async () => {
    try {
      await axios.post('http://localhost:8080/block', {
        data: newBlockData,
      });
      setNewBlockData('');
      fetchBlockchain();
      toast({
        title: 'Block added successfully',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });
    } catch (error) {
      toast({
        title: 'Error adding block',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    }
  };

  useEffect(() => {
    fetchBlockchain();
    const interval = setInterval(fetchBlockchain, 5000);
    return () => clearInterval(interval);
  }, [fetchBlockchain]);

  return (
    <Container maxW="container.lg" py={8}>
      <VStack spacing={8}>
        <Heading>Blockchain Ledger</Heading>
        <Badge colorScheme={isValid ? 'green' : 'red'} p={2} borderRadius="md">
          Chain Status: {isValid ? 'Valid' : 'Invalid'}
        </Badge>

        <Box w="100%">
          <VStack spacing={4}>
            <Input
              placeholder="Enter block data"
              value={newBlockData}
              onChange={(e) => setNewBlockData(e.target.value)}
            />
            <Button colorScheme="blue" onClick={addBlock} isDisabled={!newBlockData}>
              Add Block
            </Button>
          </VStack>
        </Box>

        <VStack w="100%" spacing={4}>
          {blocks.map((block, index) => (
            <Card key={block.hash} w="100%" variant="outline">
              <CardBody>
                <VStack align="start" spacing={2}>
                  <Text><strong>Block #{index}</strong></Text>
                  <Text><strong>Timestamp:</strong> {new Date(block.timestamp * 1000).toLocaleString()}</Text>
                  <Text><strong>Data:</strong> {block.data}</Text>
                  <Text><strong>Hash:</strong> {block.hash}</Text>
                  <Text><strong>Previous Hash:</strong> {block.prevHash || 'Genesis Block'}</Text>
                  <Text><strong>Nonce:</strong> {block.nonce}</Text>
                </VStack>
              </CardBody>
            </Card>
          ))}
        </VStack>
      </VStack>
    </Container>
  );
}

export default App;
