function onLoad() {
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
//This will ensure js is run after returning from result page
onLoad();