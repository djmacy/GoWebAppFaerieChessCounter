function onLoad() {
    updateRank1Pieces();
    updateRank2Pieces();

    window.addEventListener('load', updateRank1Pieces);
    window.addEventListener('load', updateRank2Pieces);

    document.getElementById('pawn').addEventListener('change', updateRank1Pieces);
    document.getElementById('peasant').addEventListener('change', updateRank1Pieces);
    document.getElementById('soldier').addEventListener('change', updateRank1Pieces);

    document.getElementById('rook').addEventListener('change', updateRank2Pieces);
    document.getElementById('bishop').addEventListener('change', updateRank2Pieces);
    document.getElementById('knight').addEventListener('change', updateRank2Pieces);
    document.getElementById('catapult').addEventListener('change', updateRank2Pieces);
    document.getElementById('courtesan').addEventListener('change', updateRank2Pieces);
    document.getElementById('chamberlain').addEventListener('change', updateRank2Pieces);
    document.getElementById('herald').addEventListener('change', updateRank2Pieces);
    document.getElementById('inquisitor').addEventListener('change', updateRank2Pieces);
    document.getElementById('lancer').addEventListener('change', updateRank2Pieces);
    document.getElementById('pontiff').addEventListener('change', updateRank2Pieces);
    document.getElementById('thief').addEventListener('change', updateRank2Pieces);
    document.getElementById('tower').addEventListener('change', updateRank2Pieces);
}


function saveSelections() {
    localStorage.setItem('pawn', document.getElementById('pawn').value);
    localStorage.setItem('peasant', document.getElementById('peasant').value);
    localStorage.setItem('soldier', document.getElementById('soldier').value);
    localStorage.setItem('rook', document.getElementById('rook').value);
    localStorage.setItem('bishop', document.getElementById('bishop').value);
    localStorage.setItem('knight', document.getElementById('knight').value);
    localStorage.setItem('catapult', document.getElementById('catapult').value);
    localStorage.setItem('courtesan', document.getElementById('courtesan').value);
    localStorage.setItem('chamberlain', document.getElementById('chamberlain').value);
    localStorage.setItem('herald', document.getElementById('herald').value);
    localStorage.setItem('inquisitor', document.getElementById('inquisitor').value);
    localStorage.setItem('lancer', document.getElementById('lancer').value);
    localStorage.setItem('pontiff', document.getElementById('pontiff').value);
    localStorage.setItem('thief', document.getElementById('thief').value);
    localStorage.setItem('tower', document.getElementById('tower').value);
    localStorage.setItem('king', document.getElementById('king').value);
    localStorage.setItem('queen', document.getElementById('queen').value);
    localStorage.setItem('diffLabel', document.getElementById('diffLabel').value);
}

function loadSelections() {
    document.getElementById('pawn').value = localStorage.getItem('pawn') || '4';
    document.getElementById('peasant').value = localStorage.getItem('peasant') || '0';
    document.getElementById('soldier').value = localStorage.getItem('soldier') || '0';
    document.getElementById('rook').value = localStorage.getItem('rook') || '0';
    document.getElementById('bishop').value = localStorage.getItem('bishop') || '0';
    document.getElementById('catapult').value = localStorage.getItem('catapult') || '0';
    document.getElementById('courtesan').value = localStorage.getItem('courtesan') || '0';
    document.getElementById('chamberlain').value = localStorage.getItem('chamberlain') || '0';
    document.getElementById('herald').value = localStorage.getItem('herald') || '0';
    document.getElementById('inquisitor').value = localStorage.getItem('inquisitor') || '0';
    document.getElementById('lancer').value = localStorage.getItem('lancer') || '0';
    document.getElementById('pontiff').value = localStorage.getItem('pontiff') || '0';
    document.getElementById('thief').value = localStorage.getItem('thief') || '0';
    document.getElementById('tower').value = localStorage.getItem('tower') || '0';
    document.getElementById('king').value = localStorage.getItem('king') || 'King';
    document.getElementById('queen').value = localStorage.getItem('queen') || 'Queen';
    document.getElementById('diffLabel').value = localStorage.getItem('diffLabel') || 'Beginner';
    updateRank1Pieces();
    updateRank2Pieces();
}

function updateRank1Pieces() {
        const pawnValue = parseInt(document.getElementById('pawn').value);
        const soldierValue = parseInt(document.getElementById('soldier').value);
        const peasantValue = parseInt(document.getElementById('peasant').value);

        const totalRank1Pieces = pawnValue + peasantValue + soldierValue;
        const maxRank1Pieces = 8;

        const rank1Label = document.getElementById('rank1Label');
        rank1Label.textContent = "Rank I Pieces Left: " + (maxRank1Pieces - totalRank1Pieces);
    }

function updateRank2Pieces() {
        const rookValue = parseInt(document.getElementById('rook').value);
        const bishopValue = parseInt(document.getElementById('bishop').value);
        const knightValue = parseInt(document.getElementById('knight').value);
        const chamberlainValue = parseInt(document.getElementById('chamberlain').value);
        const courtesanValue = parseInt(document.getElementById('courtesan').value);
        const catapultValue = parseInt(document.getElementById('catapult').value);
        const heraldValue = parseInt(document.getElementById('herald').value);
        const inquisitorValue = parseInt(document.getElementById('inquisitor').value);
        const lancerValue = parseInt(document.getElementById('lancer').value);
        const pontiffValue = parseInt(document.getElementById('pontiff').value);
        const thiefValue = parseInt(document.getElementById('thief').value);
        const towerValue = parseInt(document.getElementById('tower').value);

        const totalRank2Pieces = rookValue + bishopValue + knightValue + chamberlainValue + courtesanValue +
            catapultValue + heraldValue + inquisitorValue + lancerValue + pontiffValue + thiefValue + towerValue;
        const maxRank2Pieces = 6;
        const rank2Label = document.getElementById('rank2Label');
        rank2Label.textContent = "Rank II Pieces Left: " + (maxRank2Pieces - totalRank2Pieces);
    }

//window.addEventListener('load', updateRank1Pieces);
//window.addEventListener('load', updateRank2Pieces);