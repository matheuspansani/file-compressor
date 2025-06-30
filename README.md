# File-Compressor (Huffman Encoding in Go)

Este projeto tem como objetivo implementar um algoritmo de **compactação de arquivos usando Huffman**, desenvolvido em **Go (Golang)**.  
Ele permite ler arquivos binários ou texto, construir uma árvore de Huffman baseada na frequência dos bytes e gerar um novo arquivo compactado.

> ⚠️ **Status: Em desenvolvimento**

---

## Funcionalidades planejadas

- [x] Leitura de arquivos em formato `[]byte`
- [x] Geração de mapa de frequência dos bytes
- [x] Estrutura de árvore de Huffman
- [ ] Heap de nós com base na frequência
- [ ] Construção da árvore binária de Huffman
- [ ] Geração do mapa de códigos binários (`map[byte]string`)
- [ ] Codificação do conteúdo original
- [ ] Escrita do arquivo compactado (`.huff`)
- [ ] Descompactação (opcional)

---

## Conceito por trás

O algoritmo de Huffman é uma técnica de compressão **sem perdas** baseada na frequência dos símbolos em um arquivo.  
Os símbolos mais frequentes recebem códigos binários mais curtos, reduzindo o tamanho total do arquivo ao ser reescrito com esses códigos.

